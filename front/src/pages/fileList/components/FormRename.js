import { useAlert } from "../../../hooks/useAlert";
import getPathFromLocation from "../../../utils/path";
import requestSender from "../../../utils/requestSender";
import Validator from "../../../utils/validator";
import Alert from "../../layout/Alert";

function FormRename({path, name}) {
    const [alert, setAlert] = useAlert()

    const onSubmit = async (event) => {
        event.preventDefault()

        const error = (new Validator("name", event.target.name.value)).maxLength(100).required().getFirstError()
        if (error != null) {
            setAlert({type: 'warning', message: error})
        } else {
            var formData = new FormData()
            formData.append("path", getPathFromLocation(path))
            formData.append("name", event.target.name.value)
            formData.append("oldName", name)

            const res = await requestSender('/files/rename', formData, {method: 'POST', setAlert: setAlert})
            if (res.statusCode === 200) {
                setAlert({type: 'success', message: 'File renamed'})
            }
        }        
    }

    return <form onSubmit={onSubmit}>
        <div className="input-box">
            <label for="name">Name</label>
            <input id="name" type="text" name="name" max="100" defaultValue={name} autoFocus={true}/>
        </div>
        <button type="submit" className="shadow btn fix-blink">Save</button>
        <Alert alert={alert}/>
    </form>
}

export default FormRename