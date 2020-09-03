import React from "react";
import { BrowserRouter as Router, Redirect, Route, Switch } from "react-router-dom";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";

import Header from "./components/Header";
import Home from "./pages/Home";
import Dashboard from "./pages/Dashboard";
import MakeWeb from "./pages/MakeWeb";
import RegistLesson from "./pages/RegistLesson";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import AccountForm from "./pages/AccountForm";
import Website from "./pages/Website";
import useUser from "./hooks/useUser";
import ReserveForm from "./pages/ReserveForm";
import Reserved from "./pages/Reserved";
import NotFound from "./pages/NotFound";

const PrivateRoute = (props) => {
  console.log(props.user);
  if (props.user !== null) {
    return <Route exact path={props.path} component={props.component} />;
  } else {
    return <Redirect to="/" />;
  }
};

const theme = createMuiTheme({
  palette: {
    primary: {
      main: "#34675C",
    },
    secondary: {
      main: "#B3C100",
    },
  },
});

const App = () => {
  const { user, logout } = useUser();

  const withHeader = (Component: React.FC | typeof React.Component): any => {
    return () => (
      <>
        <Header user={user} logout={logout} />
        <Component />
      </>
    );
  };

  return (
    <MuiThemeProvider theme={theme}>
      <div className="container">
        <Router>
          <Switch>
            <Route exact path="/" component={withHeader(Home)} />
            <Route exact path="/login" component={withHeader(Login)} />
            <Route exact path="/signup" component={withHeader(Signup)} />
            <Route exact path="/owner/:owner_id/web" component={Website} />
            <Route exact path="/owner/:owner_id/web/:lesson_id" component={ReserveForm} />
            <Route exact path="/reserved" component={Reserved} />
            <PrivateRoute
              user={user}
              exact
              path="/regist_owner"
              component={withHeader(AccountForm)}
            />
            <PrivateRoute user={user} exact path="/dashboard" component={withHeader(Dashboard)} />
            <PrivateRoute user={user} exact path="/make_web" component={withHeader(MakeWeb)} />
            <PrivateRoute
              user={user}
              exact
              path="/regist_lesson"
              component={withHeader(RegistLesson)}
            />
            <Route component={NotFound} />
          </Switch>
        </Router>
      </div>
    </MuiThemeProvider>
  );
};

export default App;
