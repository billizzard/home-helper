import BackButton from "../BackButton";

function FileInfo({data}) {
    const backendUrl = window.appConfig.backendUrl;
    const style = {}
    if (data.PublicLink != "") {
        style.backgroundImage = `url("${backendUrl}/${data.PublicLink}")`
        style.backgroundSize = 'cover'
    }
    return <div className="fi-body">
        <div>
            <BackButton link={data.PrevPath}/>
        </div>
        <div className="fi-box">
            <div className="fi-view">
                {data.PublicLink != "" ? <img src={`${backendUrl}/${data.PublicLink}`} /> : ""}
                {data.PublicLink == "" ? <i className={"bi bi-" + data.Icon}/> : ""}
            </div>
            <div className="fi-info">
                <div className="fi-info-line">Name: {data.Title}</div>
                <div className="fi-info-line">Size: {data.Size}</div>
            </div>
        </div>
    </div>
}

export default FileInfo