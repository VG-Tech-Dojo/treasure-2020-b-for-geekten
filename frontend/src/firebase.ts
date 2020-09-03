import firebase from "firebase/app";
import "firebase/auth";

const firebaseConfig = {
  apiKey: "",
  authDomain: "",
  databaseURL: "",
  projectId: "",
  storageBucket: "",
  messagingSenderId: "",
  appId: "",
  measurementId: ""
};
// Initialize Firebase
firebase.initializeApp(firebaseConfig);

const FirebaseFactory = () => {
  let auth = firebase.auth();
  return {
    auth,

    create(email: string, password: string) {
      return auth.createUserWithEmailAndPassword(email, password);
    },

    login(email: string, password: string) {
      return auth.signInWithEmailAndPassword(email, password);
    },

    logout() {
      return auth.signOut();
    },
  };
};

export default FirebaseFactory();
