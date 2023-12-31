import { useAlert } from "../../../hooks/useAlert";
import getPathFromLocation from "../../../utils/path";
import requestSender from "../../../utils/requestSender";
import Alert from "../../layout/Alert";

function FormDelete({path, items}) {
    const [alert, setAlert] = useAlert()

    const onSubmit = async (event) => {
        event.preventDefault()
        var formData = new FormData()
        formData.append("path", getPathFromLocation(path))
        formData.append("items", JSON.stringify(items))

        const res = await requestSender('/files/delete', formData, {method: 'POST', setAlert: setAlert})
        if (res.statusCode === 200) {
            setAlert({type: 'success', message: 'Files deleted'})
        }
    }

    return <form onSubmit={onSubmit}>
        <button type="submit" className="shadow btn fix-blink">Confirm</button>
        <Alert alert={alert}/>
    </form>
}

export default FormDelete