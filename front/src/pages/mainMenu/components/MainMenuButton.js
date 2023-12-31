import { Link } from "react-router-dom"

function MainMenuButton({name, icon, link}) {
    return  <Link className={"shadow fix-blink mp-btn"} to={link}>
        <i className={"bi bi-" + icon}/>
        <div className="mp-btn_title">{name}</div>
    </Link>
}

export default MainMenuButton