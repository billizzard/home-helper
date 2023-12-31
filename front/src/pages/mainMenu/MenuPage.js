import MainMenuButton from "./components/MainMenuButton";
import { useRequest } from "../../hooks/useRequest";

function MenuPage() {
  const response = useRequest('/')

  return (
    <div className="mp-body">
      {response.data != null ?
        response.data.MenuItems.map((item) => <MainMenuButton key={item.Icon} name={item.Name} icon={item.Icon} link={item.Link}/>) :
        ""  
      }
    </div>
    
  );
}

export default MenuPage;
