import BackButton from "./BackButton"
import FileButton from "./FileButton"
import Alert from "../../layout/Alert";
import { AppContext } from "../../../context/appContext";
import { useContext } from "react";

function FileList({data, onFileSelect, selected, onChangeFolder, hasPreview}) {
    const app = useContext(AppContext);
    return  (
        data == null ? "" :
            <div className="files-box">
                {data.FilesCount == 0
                    ? <div className="body-message">Category is empty</div>
                    : ""}

                {data.PrevPath != ''
                ? <BackButton link={data.PrevPath} onChangeFolder={onChangeFolder}/>
                : ""}

                {data.Folders.map(item => 
                    <FileButton key={item.Name} item={item} type="folder" onSelect={onFileSelect} selected={selected.has(item.Name)} onChangeFolder={onChangeFolder} />
                )}

                {data.Files.map(item => 
                    <FileButton key={item.Name} item={item} onSelect={onFileSelect} selected={selected.has(item.Name)} onChangeFolder={onChangeFolder} hasPreview={app.appData.hasPreview}/>
                )}

                <Alert alert={alert}/>
            </div>
    )
}

export default FileList