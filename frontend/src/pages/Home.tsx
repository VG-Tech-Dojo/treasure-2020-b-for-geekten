import React, { useEffect } from "react";
import { Link, useHistory } from "react-router-dom";
import useUser from "./../hooks/useUser";
// @ts-ignore
import logo from '../images/logo.png';

const Home: React.FC = () => {
  const { user } = useUser();
  const history = useHistory();
  useEffect(() => {
    if (user === null) {
      return;
    }
    history.replace("/dashboard");
  }, [user]);

  const loginLinkStyle = {
    color: "black",
    textDecoration: "none",
    fontSize: 20,
    paddingRight: "10px",
  }

  const signUpLinkStyle = {
    color: "white",
    textDecoration: "none",
    fontSize: 20,
    paddingLeft: "10px",
  }

  const containerStyle = {
    textAlign: "center" as const,
    flex: 1,
    paddingTop: "10%",
  };

  const logoStyle = {
    width: "40%",
  }

  const titleStyle = {
    flex: 1,
  };

  const subTitleStyle = {
    flex: 1,
    fontSize: "20px",
    paddingTop: "3%",
    fontWeight: "bold" as const,
  };

  const buttonContainerStyle = {
    paddingTop: "3%",
    flex: 1,
    display: "flex",
    flexDirection: "row" as const,
    justifyContent: "center",
    margin: "30px",
  };

  const signUpButtonStyle = {
    backgroundColor: "#B3C100",
    width: "150px",
    border: "2px solid #ffffff",
    borderRadius: "50px",
    height: "64px",
  }

  const loginButtonStyle = {
    backgroundColor: "#fff",
    width: "150px",
    border: "2px solid #34675C",
    borderRadius: "50px",
    height: "64px",
  }

  return (
    <div style={containerStyle}>
      <div style={titleStyle}>
        <img style={logoStyle} src={logo}></img>
      </div>
      <div style={subTitleStyle}>
        <p>オンラインレッスン支援サービス SlimLine</p>
        <p>あなたのレッスンを、気軽にオンラインで。</p>
      </div>
      <div style={buttonContainerStyle}>
          <Link style={loginLinkStyle} to="/login">
            <button style={loginButtonStyle}>
              ログイン
            </button>
          </Link>
          <Link style={signUpLinkStyle} to="/signup">
            <button style={signUpButtonStyle}>
            新規登録
            </button>
          </Link>
      </div>
    </div>
  );
};

export default Home;
