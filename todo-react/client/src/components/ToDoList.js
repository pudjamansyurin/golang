import React, { useState, useEffect } from "react";
import { Card, Form, Header, Input } from "semantic-ui-react";
import * as api from "../api/task";
import TheButton from "./TheButton";

export default function ToDoList() {
  const [task, setTask] = useState("");
  const [items, setItems] = useState([]);

  const doTask = (fn, { _id }) => {
    fn(_id).then(() => getTask());
  };

  const getTask = () => {
    api.getTask().then((res) => {
      if (res.data) {
        setItems(
          res.data.map((item) => {
            let color = "yellow";
            let style = {
              wordWrap: "break-word",
            };

            if (item.status) {
              color = "green";
              style["textDecorationLine"] = "line-through";
            }

            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={style}>{item.task}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <TheButton
                      icon="check circle"
                      color="green"
                      name="Done"
                      onClick={() => doTask(api.updateTask, item)}
                    ></TheButton>
                    <TheButton
                      icon="undo"
                      color="yellow"
                      name="Undo"
                      onClick={() => doTask(api.undoTask, item)}
                    ></TheButton>
                    <TheButton
                      icon="delete"
                      color="red"
                      name="Delete"
                      onClick={() => doTask(api.deleteTask, item)}
                    ></TheButton>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          })
        );
      } else setItems([]);
    });
  };

  const onSubmit = (e) => {
    if (!task) return;

    api.createTask(task).then((res) => {
      getTask();
      setTask("");
    });
  };

  useEffect(() => {
    getTask();

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <React.Fragment>
      <div className="row">
        <Header className="header" as="h2">
          TODO LIST
        </Header>
      </div>
      <div className="row">
        <Form onSubmit={onSubmit}>
          <Input
            type="text"
            name="task"
            placeholder="Task name"
            onChange={(e) => setTask(e.target.value)}
            value={task}
            fluid
          />
        </Form>
      </div>
      <div className="row">
        <Card.Group>{items}</Card.Group>
      </div>
    </React.Fragment>
  );
}
