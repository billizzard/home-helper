import getPathFromLocation from "../../../utils/path";
import requestSender from "../../../utils/requestSender";

function ActionDownloadButton({path, items}) {
    const addButtonClick = async (event) => {
        event.preventDefault()
        var formData = new FormData()
        const itemsArr = []

        items.forEach((value, key, map) => {
            itemsArr.push(key)
        });
        formData.append("path", getPathFromLocation(path))
        formData.append("items", JSON.stringify(itemsArr))

        const res = await requestSender('/files/download', formData, {method: 'POST', isFile: true})
    }

    return <div>
        <a className={"shadow files_action_btn fix-blink"} onClick={addButtonClick} download="File">
            <i className={"bi bi-upload"}/>
        </a>
    </div>
}

export default ActionDownloadButton