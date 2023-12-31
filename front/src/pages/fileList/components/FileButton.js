import { Link } from "react-router-dom"

function FileButton({item, type = '', onSelect, selected = false, onChangeFolder, hasPreview}) {
    const backendUrl = window.appConfig.backendUrl;
    const onClick = () => {
        onSelect(item.Name)
    }

    const showImage = (item.PublicLink !== "") && hasPreview
    const style = {}

    if (showImage) {
        style.backgroundImage = `url("${backendUrl}/${item.PublicLink}")`
        style.backgroundSize = 'cover'
    }

    return  <div className={"file_box "}>
        <Link className={"shadow file-btn fix-blink " + type + (selected ? " selected" : "")} onClick={onChangeFolder} to={item.Link} style={style}>
            {showImage ? "" : <i className={"bi bi-" + item.Icon}/>}
        </Link>
        <div className="file_desc" onClick={onClick}>{item.Name}</div>
    </div>
}

export default FileButton