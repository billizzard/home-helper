import { useRef } from "react";
import { useAlert } from "../../../hooks/useAlert";
import getPathFromLocation from "../../../utils/path";
import requestSender from "../../../utils/requestSender";
import Alert from "../../layout/Alert";

function ActionListUploadButton({path}) {
    const [alert, setAlert] = useAlert()

    const inputRef = useRef(null);

    const handleClick = () => {
        inputRef.current.click();
    };

    const handleFileChange = async (event) => {
        const files = event.target.files
        const fileObj = event.target.files && event.target.files[0];
        if (!fileObj) {
            return;
        }

        var formData = new FormData()
        for (var x = 0; x < files.length; x++) {
            formData.append("fileuploadfiles", files[x]);
        }

        formData.append("path", getPathFromLocation(path))
        event.target.value = null;

        const res = await requestSender('/files/upload', formData, {method: 'POST', setAlert: setAlert})

        if (res.statusCode === 200) {
            setAlert({type: 'success', message: 'Files uploaded successfully'})
        }
    };

    return <div className={"shadow files_action_btn fix-blink"} onClick={handleClick}>
        <input
            style={{display: 'none'}}
            ref={inputRef}
            type="file"
            multiple = {true}
            onChange={handleFileChange}
            name="uploadfiles"
        />
        <i className={"bi bi-download"}/>
        <Alert alert={alert}/>
    </div>
}

export default ActionListUploadButton