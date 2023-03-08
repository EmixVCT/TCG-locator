import React, { useState } from "react";
import { ScrollArea } from "@/components/ui/scroll-area"
import { Icons } from "@/components/icons"

const getDefaultSorting = (defaultTableData, columns) =>{
  const sorted = [...defaultTableData].sort((a, b) => {
    const filterColumn = columns.filter((column) => column.sortbyOrder);

    // Merge all array objects into single object and extract accessor and sortbyOrder keys
    let { accessor = "id", sortbyOrder = "asc" } = Object.assign(
      {},
      ...filterColumn
    );

    if (a[accessor] === null) return 1;
    if (b[accessor] === null) return -1;
    if (a[accessor] === null && b[accessor] === null) return 0;

    const ascending = a[accessor]?
      .toString()
      .localeCompare(b[accessor].toString(), "en", {
        numeric: true,
      });

    return sortbyOrder === "asc" ? ascending : -ascending;
  });
  return sorted;
}

const useSortableTable = (data, columns) => {
  const [tableData, setTableData] = useState(getDefaultSorting(data, columns));

  const handleSorting = (sortField, sortOrder) => {
    if (sortField) {
      const sorted = [...tableData].sort((a, b) => {
        if (a[sortField] === null) return 1;
        if (b[sortField] === null) return -1;
        if (a[sortField] === null && b[sortField] === null) return 0;
        return (
          a[sortField].toString().localeCompare(b[sortField].toString(), "en", {
            numeric: true,
          }) * (sortOrder === "asc" ? 1 : -1)
        );
      });
      setTableData(sorted);
    }
  };

  return [tableData, handleSorting];
};
const TableHead = ({ columns, handleSorting }) => {
  const [sortField, setSortField] = useState("");
  const [order, setOrder] = useState("asc");

  const handleSortingChange = (accessor) => {
    const sortOrder =
      accessor === sortField && order === "asc" ? "desc" : "asc";
    setSortField(accessor);
    setOrder(sortOrder);
    handleSorting(accessor, sortOrder);
  };

  return (
    <thead className="">
      <tr>
        {columns.map(({ label, accessor, sortable }) => {
          return (
            <th
              key={accessor}
              onClick={sortable ? () => handleSortingChange(accessor) : null}
              className={`border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-slate-400 dark:text-slate-200 text-left ${sortable && "cursor-pointer"}`}
            >
              <div className="flex items-center justify-between">
                {label}
                <span className="ml-4">
                  {sortable
                    ? sortField === accessor && order === "asc"
                      ? <Icons.sortasc className="h-4 w-4" />
                      : sortField === accessor && order === "desc"
                      ? <Icons.sortdesc className="h-4 w-4" />
                      : <Icons.sort className="h-4 w-4" />
                    : ""}
                </span>
              </div>
            </th>
          );
        })}
      </tr>
    </thead>
  );
};

const TableBody = ({ tableData, columns }) => {
  return (
    <tbody className="bg-white dark:bg-slate-800">
      {tableData.map((data) => {
        return (
          <tr key={data.id}>
            {columns.map(({ accessor, isBoolean }) => {
              const tData = data[accessor] ? data[accessor] : "——";
              return <td key={accessor} className="border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-slate-500 dark:text-slate-100">
                {isBoolean
                  ? tData ? <Icons.yes className="h-4 w-4 text-green-400" /> : <Icons.no className="h-4 w-4 text-red-400" />
                  : tData}
              </td>;
            })}
          </tr>
        );
      })}
    </tbody>
  );
};
const Table = ({ data, columns }) => {
  const [tableData, handleSorting] = useSortableTable(data, columns);

  return (
    <div className="not-prose relative bg-slate-50 rounded-xl overflow-hidden dark:bg-slate-800/25 relative rounded-xl border border-gray-200 dark:border-gray-800">
      <div className="overflow-x-auto my-8">
        <table className="border-collapse table-auto w-full text-sm">
          <TableHead {...{ columns, handleSorting }} />
          <TableBody {...{ columns, tableData }} />
        </table>
      </div>
    </div>
  );
};

export default Table;
