import "./App.css";
import "semantic-ui-css/semantic.min.css";

import { Container } from "semantic-ui-react";
import ToDoList from "./components/ToDoList";

function App() {
  return (
    <Container>
      <ToDoList />
    </Container>
  );
}

export default App;
