import { Routes, Route, BrowserRouter } from "react-router-dom";
import Navbar from "./components/Navbar";
import Profile from "./pages/ProfilePage";
import PostsPage from "./pages/PostsPage";
import Post from "./components/Post";
import Signup from "./pages/SignupPage";
import Login from "./pages/LoginPage";
import NoPage from "./pages/NoPage";
import Category from "./pages/CategoryPage";
import RequireAuth from "./components/RequireAuth";
import ProfileInfo from "./components/ProfileInfo";
import GroupPage from "./pages/GroupPage";
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
              <Route path="profile/:id" element={<ProfileInfo />} />
              <Route
                path="posts"
                element={<PostsPage showCreatePost={true} />}
              />
              <Route path="groups/:groupId" element={<GroupPage />} />
              <Route path="posts/:id" element={<Post />} />
              <Route path="categories/:id" element={<Category />} />
              <Route path="logout" element={<Logout />} />
              <Route path="groups/:groupId" element={<GroupPage />} />
              <Route path="*" element={<NoPage />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
