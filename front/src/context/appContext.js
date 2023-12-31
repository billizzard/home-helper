import { createContext, useState } from 'react';

export const AppContext = createContext({
    hasPreview: false
  });

  export const AppContextProvider = ({children, initial = {
    hasPreview: localStorage.getItem('imagePreview') !== 'false',
  }}) => {
    const [appData, setAppData] = useState(initial)

    return <AppContext.Provider value={{appData, setAppData}}>
        {children}
    </AppContext.Provider>

  }