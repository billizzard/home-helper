import React from "react"
import Header from "./Header"
import Footer from "./Footer"
import { AppContextProvider } from "../../context/appContext"
function Layout(props)  {
  return (
    <AppContextProvider>
      <div className="body">
        <Header />
          {props.children}
        <Footer />
      </div>
    </AppContextProvider>
  )
}
export default Layout;