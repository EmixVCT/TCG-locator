import * as React from "react"
import Table from "@/components/ui/table"
import { ScrollArea } from "@/components/ui/scroll-area"

import data from "@/data.js"

/*
{"Url":"https://shop.fantasysphere.net/one-piece-tcg/10024012--pre-commande-one-piece-display-boite-de-24-boosters-paramount-war-op-02-en-811039039097.html","Label":"One piece TCG Display OP02","Country":"FR","Language":"ENG","Currency":"EUR","State":true,"LastFetchDate":"2023-03-06T16:14:05.016467902Z","LastStockDate":"0001-01-01T00:00:00Z","Stock":false,"Price":"90"}
*/
export const DealsTable = () => {
  const columns = [
    { label: "Item", accessor: "Label", sortable: true },
    { label: "In Stock", accessor: "Stock", isBoolean: true, sortable: true, sortbyOrder: "desc" },
    { label: "Price", accessor: "Price", sortable: true },
    { label: "Updateed", accessor: "LastFetchDate", sortable: false},
  ];
  const d = data.map((d, i) => ({...d, id: i}))

  return (
    <div className="container">
      <Table columns={columns} data={d} />
    </div>
  )
}
