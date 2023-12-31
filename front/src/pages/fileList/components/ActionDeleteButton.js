import { useState } from "react";
import Modal from "../../layout/Modal";
import FormDelete from "./FormDelete";

function ActionDeleteButton({path, items}) {
    const [isShow, setShow] = useState(false)

    const addButtonClick = () => {
        setShow(true)
    }

    const itemsArr = []

    items.forEach((value, key, map) => {
        itemsArr.push(key)
    });

    return <div>
        <div className={"shadow files_action_btn fix-blink"} onClick={addButtonClick}>
            <i className={"bi bi-trash3"}/>
        </div>
        <Modal isShow={isShow} setIsShow={setShow} title="Confirm deletion"><FormDelete path={path} items={itemsArr}/></Modal>
    </div>
}

export default ActionDeleteButton