import React from "react";
import App from "./src/App";
import { render } from "react-dom";
import { initEnv } from "./env";

const indexStyle = {
  height: "100vh",
  backgroundColor: "#FFFDF3",
  margin: "-10px",
};

initEnv();

render(
  <div style={indexStyle}>
    <App />
  </div>,
  document.getElementById("root")!
);
