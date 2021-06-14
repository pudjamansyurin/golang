import "./App.css";
import { Admin, ListGuesser, Resource } from "react-admin";
import jsonServiceProvider from "ra-data-json-server";
import { UserList } from "./components/resources/users";
import { PostCreate, PostEdit, PostList } from "./components/resources/posts";
import PostIcon from "@material-ui/icons/Book";
import UserIcon from "@material-ui/icons/Group";
import authProvider from "./providers/authProvider";
import Dashboard from "./components/dashboard";

const dataProvider = jsonServiceProvider("http://localhost:8010");

function App() {
  return (
    <Admin
      dashboard={Dashboard}
      dataProvider={dataProvider}
      authProvider={authProvider}
    >
      <Resource name="products" list={ListGuesser} />
      {/* <Resource
        name="posts"
        list={PostList}
        edit={PostEdit}
        create={PostCreate}
        icon={PostIcon}
      ></Resource>
      <Resource name="users" list={UserList} icon={UserIcon}></Resource> */}
    </Admin>
  );
}

export default App;
