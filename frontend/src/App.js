import { Routes, Route, BrowserRouter } from "react-router-dom";
import Navbar from "./components/Navbar";
import Profile from "./components/Profile";
import PostsPage from "./components/PostsPage";
import Chat from "./components/Chat";
import Post from "./components/Post";
import Signup from "./components/Signup";
import Signin from "./components/Signin";
import NoPage from "./components/NoPage";
import Category from "./components/Category";
import RequireAuth from "./components/RequireAuth";
//import RequireGuest from "./components/RequireGuest";
import Logout from "./components/Logout";

import "./style.css";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Navbar />}>
            <Route path="signin" element={<Signin />} />
            <Route path="signup" element={<Signup />} />

            <Route path="profile" element={<Profile />} />
            <Route path="posts" element={<PostsPage />} />
            <Route path="chat" element={<Chat />} />
            <Route path="posts/:id" element={<Post />} />
            <Route path="categories/:id" element={<Category />} />
            <Route path="logout" element={<Logout />} />
            <Route path="*" element={<NoPage />} />
            {/* <Route element={<RequireAuth />}></Route> */}
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
