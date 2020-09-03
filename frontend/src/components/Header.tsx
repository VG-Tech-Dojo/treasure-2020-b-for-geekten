import React from "react";
import { Link } from "react-router-dom";
import { AppBar, Toolbar, Typography, Button } from "@material-ui/core";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";

interface Props {
  user: firebase.User | null;
  logout: () => Promise<void>;
}

const linkStyle = {
  color: "#fff",
  textDecoration: "none",
  marginLeft: "30px",
  fontWeight: "bold" as const,
};

const logoutStyle = {
  marginLeft: "100px",
  color: "#fff",
  textDecoration: "none",
  fontWeight: "bold" as const,
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  })
);

const Header: React.FC<Props> = (props) => {
  const { user, logout } = props;
  const classes = useStyles();

  const logoutHandler = async () => {
    await logout();
    window.location.href = "/";
  };

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            {user !== null && (
              <Link style={linkStyle} to="/dashboard">
                SlimLine
              </Link>
            )}
            {user === null && <Typography>SlimLine</Typography>}
          </Typography>
          {user !== null && (
            <Button color="inherit">
              <Link style={linkStyle} to="/regist_owner">
                講師登録
              </Link>
            </Button>
          )}
          {user !== null && (
            <Button color="inherit">
              <Link style={linkStyle} to="/regist_lesson">
                レッスン登録
              </Link>
            </Button>
          )}
          {user !== null && (
            <Button color="inherit">
              <a style={logoutStyle} href="#" onClick={logoutHandler}>
                ログアウト
              </a>
            </Button>
          )}
        </Toolbar>
      </AppBar>
    </div>
  );
};

export default Header;
