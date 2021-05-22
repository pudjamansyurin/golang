import React from "react";
import { Button, Icon } from "semantic-ui-react";

export default function TheButton({ color, icon, name, onClick }) {
  return (
    <Button size="mini" color={color} onClick={() => onClick()} basic>
      <Icon name={icon} /> {name}
    </Button>
  );
}
