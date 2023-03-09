import { NavItem } from "@/types/nav"

interface SiteConfig {
  name: string
  description: string
  mainNav: NavItem[]
  links: {
    twitter: string
    github: string
    docs: string
  }
  apiEndpoint: string
}

export const siteConfig: SiteConfig = {
  name: "TCG Locator",
  description:
    "Find Trading Cards in Stock",
  mainNav: [
    {
      title: "Home",
      href: "/",
    },
  ],
  links: {
    twitter: "https://twitter.com/EmixVCT",
    github: "https://github.com/EmixVCT/TCG-locator",
    docs: "",
  },
  apiEndpoint: "https://api.tcglocator.xyz/locators"
}
