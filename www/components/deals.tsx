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
    { label: "Description", accessor: "Label",info: "Language", sortable: true },
    { label: "Live", accessor: "State", isBoolean: true, sortable: true },
    { label: "Vendor", accessor: "Vendor",info: "Country",sortable: true },
    { label: "Stock", accessor: "Stock", isBoolean: true, sortable: true, sortbyOrder: "desc" },
    { label: "Last Stock", accessor: "LastStockDate",date:true, sortable: true },
    { label: "Price", accessor: "Price", info:"Currency", sortable: true },

  ];
  useEffect(() => {
    const data = fetch(siteConfig.apiEndpoint)
    .then(r => r.json())
    .then(res => res.data)
    .then(data => data.map((d, i) => ({...d, id: i})))
    .then(data => setData(data))
    .catch(err => {
      console.log(err)
    })
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
