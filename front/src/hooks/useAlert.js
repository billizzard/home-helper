import { useEffect, useState } from "react";

export function useAlert(message = null, type = 'warning') {
    const [alert, setAlert] = useState({type: type, message: message});

    useEffect(() => {
    const timer = null;
        if (alert.message != null) {
            const timer = setTimeout(() => {
            setAlert({type: null, message: null})
            }, 5000);
        }
        return () => clearTimeout(timer);
    }, [alert.message])

    return [alert, setAlert];
  }