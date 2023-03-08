import React, { useState, useEffect } from "react"
import Table from "@/components/ui/table"
import { ScrollArea } from "@/components/ui/scroll-area"

import data from "@/data.js"
import { siteConfig } from "@/config/site"

/*
{"Url":"https://shop.fantasysphere.net/one-piece-tcg/10024012--pre-commande-one-piece-display-boite-de-24-boosters-paramount-war-op-02-en-811039039097.html","Label":"One piece TCG Display OP02","Country":"FR","Language":"ENG","Currency":"EUR","State":true,"LastFetchDate":"2023-03-06T16:14:05.016467902Z","LastStockDate":"0001-01-01T00:00:00Z","Stock":false,"Price":"90"}
*/

export const DealsTable = () => {
  const [data, setData] = useState(false)
  const columns = [
    { label: "Item", accessor: "Label", sortable: true },
    { label: "In Stock", accessor: "Stock", isBoolean: true, sortable: true, sortbyOrder: "desc" },
    { label: "Price", accessor: "Price", sortable: true },
    { label: "Updated", accessor: "LastFetchDate", sortable: false},
  ];
  useEffect(() => {
    const data = fetch(siteConfig.apiEndpoint)
    .then(r => r.json())
    .then(res => res.data)
    .then(data => data.map((d, i) => ({...d, id: i})))
    .then(data => setData(data))
  }, [])

  return (
    <div className="container">
      
      {data
        ? <Table columns={columns} data={data} />
        : <span>Loading</span>
      }
    </div>
  )
}
