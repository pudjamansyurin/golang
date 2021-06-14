import React from "react";

import { Datagrid, EmailField, List, TextField } from "ra-ui-materialui";
import { MyUrlField } from "../fields/MyUrlField";

export const UserList = (props) => (
  <List {...props}>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <TextField source="name" />
      <EmailField source="email" />
      <TextField source="phone" />
      <MyUrlField source="website" />
      <TextField source="company.name" />
    </Datagrid>
  </List>
);
