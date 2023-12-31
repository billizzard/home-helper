import Alert from "../../layout/Alert";
import { useState } from "react";
import FileList from "./FileList";
import ActionList from "./ActionList";
import { useLocation } from "react-router-dom";
import { useAlert } from "../../../hooks/useAlert";
import { useRequest } from "../../../hooks/useRequest";
import FileInfo from "./info/FileInfo";

function FileListBody() {
    const location = useLocation()
    const [alert, setAlert] = useAlert()
    const response = useRequest(location.pathname, null, {setAlert: setAlert})
    const data = response.data
    const [selected, setSelected] = useState(new Map())
    const [reload, setReload] = useState(1)
    const hasPreview = localStorage.getItem('imagePreview')

    const onSelect = (name) => {
        let newSelected = new Map(selected)
        setReload(2)
        if (newSelected.has(name)) {
            newSelected.delete(name)
        } else {
            newSelected.set(name, true)
        }

        setSelected(newSelected)
    }

    const onChangeFolder = () => {
        setSelected(new Map())
    }

    return (data && data.IsFileInfo) ?
    <FileInfo data={data}></FileInfo>
    :
    <>
        <ActionList path={location.pathname} selected={selected}/>
        <FileList data={data} onFileSelect={onSelect} selected={selected} onChangeFolder={onChangeFolder} hasPreview={hasPreview}/>
        <Alert alert={alert}/>
    </>
}

export default FileListBody