import { Routes, Route, BrowserRouter } from "react-router-dom";
import Navbar from "./components/Navbar";
import Profile from "./pages/ProfilePage";
import PostsPage from "./pages/PostsPage";
import Chat from "./pages/ChatPage";
import Post from "./components/Post";
import Signup from "./pages/SignupPage";
import Login from "./pages/LoginPage";
import NoPage from "./pages/NoPage";
import Category from "./pages/CategoryPage";
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
            <Route path="login" element={<Login />} />
            <Route path="signup" element={<Signup />} />

            <Route element={<RequireAuth />}>
              <Route path="profile" element={<Profile />} />
              <Route path="posts" element={<PostsPage />} />
              <Route path="chat" element={<Chat />} />
              <Route path="posts/:id" element={<Post />} />
              <Route path="categories/:id" element={<Category />} />
              <Route path="logout" element={<Logout />} />
              <Route path="*" element={<NoPage />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
