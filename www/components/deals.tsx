import React, { useState, useEffect } from "react"
import Table from "@/components/ui/table"
import { ScrollArea } from "@/components/ui/scroll-area"

import data from "@/data.js"
import { siteConfig } from "@/config/site"

/*
{"Url":"https://shop.fantasysphere.net/one-piece-tcg/10024012--pre-commande-one-piece-display-boite-de-24-boosters-paramount-war-op-02-en-811039039097.html","Label":"One piece TCG Display OP02","Country":"FR","Language":"ENG","Currency":"EUR","State":true,"LastFetchDate":"2023-03-06T16:14:05.016467902Z","LastStockDate":"0001-01-01T00:00:00Z","Stock":false,"Price":"90"}
*/
interface Locator {
  Country: String;
  Currency: String;
  Label: String;
  Language: String;
  LastFetchDate: String;
  LastStockDate: String;
  Price: String;
  State: boolean;
  Stock: boolean;
  Url: String;
  Vendor: String;
  id: number;
}

export const DealsTable = ({filter}) => {
  const [data, setData] = useState<Locator[]>()
  const [dataFiltered, setDataFiltered] = useState<Locator[]>()

  const refreshInterval = 30000
  const columns = [
    { label: "Description", accessor: "Label",info: "Language", sortable: true },
    { label: "Live", accessor: "State", isBoolean: true, sortable: true },
    { label: "Vendor", accessor: "Vendor",info: "Country",sortable: true },
    { label: "Stock", accessor: "Stock", isBoolean: true, sortable: true, sortbyOrder: "desc" },
    { label: "Last Stock", accessor: "LastStockDate",date:true, sortable: true },
    { label: "Price", accessor: "Price", info:"Currency", sortable: true },

  ];

  const applyFilter = (d,s) => {
    let filterData = d.filter((column) => column.Label.toUpperCase().includes(s.toUpperCase()))
    return filterData
  }

  const fetchData = () => {
    fetch(siteConfig.apiEndpoint)
      .then(r => r.json())
      .then(res => res.data)
      .then(d => d.map((d, i) => ({...d, id: i})))
      .then(d => setData(d))
      .catch(err => console.log(err))
  }

  useEffect(() => {
    fetchData()
    const interval = setInterval(fetchData, refreshInterval);
    return () => clearInterval(interval);
  }, [])


  useEffect(() => {
    if (data) {
      setDataFiltered(applyFilter(data,filter))
    }
  }, [data])

  useEffect(() => {
    if (data){
      setDataFiltered(applyFilter(data,filter))
    }
  }, [filter])

  return (
    <div className="container">
      
      {data && dataFiltered
        ? <Table columns={columns} data={dataFiltered} />
        : <span>Loading</span>
      }
    </div>
  )
}
