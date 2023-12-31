import { useEffect, useState } from "react";
import requestSender from "../utils/requestSender";

export function useRequest(url, data, options) {
    const settings = {
        method: options?.method || 'GET',
        setAlert: options?.setAlert || null
    }
    const [response, setResponse] = useState({isLoading: true, response: null});

    useEffect(() => {
        
        (async () => {
            const senderResponse = await requestSender(url, data, settings)
            
            if (senderResponse.redirect != null) {
                window.location.href = senderResponse.redirect
                return
            }

            if (senderResponse.error != null) {
                setResponse({
                    isLoading: false, 
                    data: null
                })

                return
            }

            setResponse({
                isLoading: false,
                data: senderResponse.data
            })
          })();
      }, [url]);

      return response;
  }