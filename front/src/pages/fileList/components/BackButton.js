import { Link } from "react-router-dom"

function BackButton({link, onChangeFolder}) {
    return  <div className="file_box">
        <Link className={"shadow file_back_btn fix-blink"} to={'/files/' + link} onClick={onChangeFolder}>
            <i className={"bi bi-arrow-left-short"}/>
        </Link>
    </div>
}

export default BackButton