import { useAlert } from "../../../hooks/useAlert";
import getPathFromLocation from "../../../utils/path";
import requestSender from "../../../utils/requestSender";
import Validator from "../../../utils/validator";
import Alert from "../../layout/Alert";

function FormAddFolder({path}) {
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

            const res = await requestSender('/files/folder/create', formData, {method: 'POST', setAlert: setAlert})
            if (res.statusCode === 200) {
                setAlert({type: 'success', message: 'Folder created'})
            }
        }        
    }

    return <form onSubmit={onSubmit}>
        <div className="input-box">
            <label for="name">Name</label>
            <input id="name" type="text" name="name" max="100" autoFocus={true}/>
        </div>
        <button type="submit" className="shadow btn fix-blink">Create</button>
        <Alert alert={alert}/>
    </form>
}

export default FormAddFolder