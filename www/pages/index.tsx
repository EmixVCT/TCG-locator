import Head from "next/head"
import Link from "next/link"

import { siteConfig } from "@/config/site"
import { Layout } from "@/components/layout"
import { buttonVariants } from "@/components/ui/button"
import React, { useState, useEffect } from "react"

import { DealsTable } from "@/components/deals"
import { Input } from "@/components/ui/input"

export default function IndexPage() {
  const [searchFilter, setSearchFilter] = useState("");

  return (
    <Layout>
      <Head>
        <title>TCG Locator</title>
        <meta name="description" content="Find trading cards in stock, through the web" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <section className="container grid items-center gap-6 pt-6 pb-6 md:py-8">
        <div className="flex max-w-[980px] flex-col items-start gap-2">
          <h1 className="text-3xl font-extrabold leading-tight tracking-tighter sm:text-3xl md:text-5xl lg:text-6xl"> 
            Find trading cards in stock <br className="hidden sm:inline" />
            through the web
          </h1>
          <p className="max-w-[700px] text-lg text-slate-700 dark:text-slate-400 sm:text-xl">
            And be notified for the next deals
          </p>
        </div>
        <div className="flex gap-4">
          <Link
            href={siteConfig.links.docs}
            target="_blank"
            rel="noreferrer"
            className={buttonVariants({ size: "lg" })}
          >
            Get notified
          </Link>
          <Link
            target="_blank"
            rel="noreferrer"
            href={siteConfig.links.github}
            className={buttonVariants({ variant: "outline", size: "lg" })}
          >
            Support the project
          </Link>
        </div>
        <div className="flex gap-4 ">
          <h3 className="my-auto">Search: </h3>
          <Input 
            placeholder="Filter"
            onChange={event => setSearchFilter(event.target.value)}
          />
        </div>
      </section>
      <section className="container ">
        <div className="flex max-w-[980px] items-start gap-2 text-green-400 mb-2">
          <span className="relative flex h-3 w-3 my-auto">
            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
            <span className="relative inline-flex rounded-full h-3 w-3 bg-green-500"></span>
          </span>
          Live 
        </div>
      </section>
      <section>
        <DealsTable filter={searchFilter}/>
      </section>
    </Layout>
  )
}
