import { useState } from "react";
import Modal from "../../layout/Modal";
import FormAddFolder from "./FormAddFolder";

function ActionCreateFolderButton({path}) {
    const [isShow, setShow] = useState(false)

    const addButtonClick = () => {
        setShow(true)
    }

    return <div>
        <div className={"shadow files_action_btn fix-blink"} onClick={addButtonClick}>
            <i className={"bi bi-folder-plus"}/>
        </div>
        <Modal isShow={isShow} setIsShow={setShow} title="Enter folder name"><FormAddFolder path={path}/></Modal>
    </div>
}

export default ActionCreateFolderButton