import { useState } from "react";
import Modal from "../../layout/Modal";
import FormRename from "./FormRename";

function ActionRenameButton({path, name}) {
    const [isShow, setShow] = useState(false)

    const addButtonClick = () => {
        setShow(true)
    }

    return <div>
        <div className={"shadow files_action_btn fix-blink"} onClick={addButtonClick}>
            <i className={"bi bi-pencil"}/>
        </div>
        <Modal isShow={isShow} setIsShow={setShow} title="Edit name"><FormRename path={path} name={name}/></Modal>
    </div>
}

export default ActionRenameButton