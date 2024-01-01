import ActionCreateFolderButton from "./ActionCreateFolderButton"
import ActionDeleteButton from "./ActionDeleteButton"
import ActionDownloadButton from "./ActionDownloadButton"
import ActionListUploadButton from "./ActionListUploadButton"
import ActionRenameButton from "./ActionRenameButton"

function ActionList({path, selected}) {
    const getSelectedButtons = () => {
        const res = []
        if (selected.size === 0) {
            return res
        }

        res.push(<ActionDeleteButton path={path} items={selected}/>)
        res.push(<ActionDownloadButton path={path} items={selected}/>)

        if (selected.size === 1) {
            selected.forEach((value, key, map) => {
                res.push(<ActionRenameButton path={path} name={key}/>)
            })
        }

        return res
    }

    return  <div className="action-box">
        <div className="action-subbox">
            {getSelectedButtons()}
        </div>

        <div className="action-subbox">
            {selected.size > 0 ? <div className="action-separator"></div> : ""}
            <ActionListUploadButton path={path} />
            <ActionCreateFolderButton path={path}/>
        </div>
    </div>
}

export default ActionList