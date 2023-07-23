// eslint-disable-next-line no-unused-vars
import React from "react";

const App = () => {
  return (
    <>
      <h1>Hello word!</h1>
      <button
        onClick={async () => {
          const resp = await fetch("/user");
          const data = await resp.json();
          console.log(data);
        }}
      >
        HERE
      </button>
    </>
  );
};

export default App;
