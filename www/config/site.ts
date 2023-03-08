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
  name: "TGS Locator",
  description:
    "Find Trading Cards in Stock",
  mainNav: [
    {
      title: "Home",
      href: "/",
    },
  ],
  links: {
    twitter: "https://twitter.com/TGSLocator",
    github: "https://github.com/EmixVCT/TCG-locator",
    docs: "#",
  },
  apiEndpoint: "http://maximevincent.fr:3050/locators"
}
