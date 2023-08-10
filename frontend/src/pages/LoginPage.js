import React from "react";
import { useState, useEffect } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from "axios";
import useAuth from "../hooks/useAuth";
import Button from "react-bootstrap/Button";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Form from "react-bootstrap/Form";
import Container from "react-bootstrap/Container";
import FloatingLabel from "react-bootstrap/FloatingLabel";
import Alert from "react-bootstrap/Alert";
import { LinkContainer } from "react-router-bootstrap";

const LOGIN_URL = "http://localhost:8000/login";

const Login = () => {
  const { setAuth } = useAuth();

  const navigate = useNavigate();
  const location = useLocation();
  const from =
    location.state?.from?.pathname !== "/logout"
      ? location.state?.from?.pathname || "/profile"
      : "/profile";

  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const [errMsg, setErrMsg] = useState("");

  console.log(from, "FROM LOGIN");

  const handleChange = (event) => {
    const { name, value } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: value,
      };
    });
  };

  useEffect(() => {
    setErrMsg("");
  }, [formData]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post(LOGIN_URL, JSON.stringify(formData), {
        headers: { "Content-Type": "application/json" },
        withCredentials: true,
      });

      console.log(JSON.stringify(response, null, 2));
      setAuth(true);
      setFormData({
        username: "",
        password: "",
      });

      navigate(from, { replace: true });
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else if (err.response?.status === 400) {
        setErrMsg("Missing username or password");
      } else if (err.response?.status === 401) {
        setErrMsg("Wrong username or password");
      } else {
        setErrMsg("Login Failed");
      }
    }
  };

  return (
    <Container>
      <Row className="justify-content-center">
        <Col sm="6" className="text-center">
          {errMsg && <Alert variant="danger">{errMsg}</Alert>}
        </Col>
      </Row>

      {/* <Row className="justify-content-md-center">
        <Col xs lg="2">
          1 of 3
        </Col>
        <Col md="1">Variable width content</Col>
        <Col xs lg="2">
          3 of 3
        </Col>
      </Row> */}
      <Row className="justify-content-center">
        <Col sm="6" className="border rounded p-3">
          <Form onSubmit={handleSubmit}>
            <FloatingLabel
              className="mb-3"
              controlId="floatingEmail"
              label="Email address or username"
            >
              <Form.Control
                type="email"
                placeholder="Email address"
                onChange={handleChange}
                name="username"
                value={formData.username}
                required
                autoFocus
              />
            </FloatingLabel>
            <FloatingLabel
              controlId="floatingPassword"
              className="mb-3"
              label="Password"
            >
              <Form.Control
                type="password"
                placeholder="Password"
                onChange={handleChange}
                name="password"
                value={formData.password}
                required
              />
            </FloatingLabel>
            <Col as={Button} xs="12" type="submit">
              Sign In
            </Col>
          </Form>
        </Col>
      </Row>

      <Row className="justify-content-center">
        <Col sm="6" className="text-center mt-3">
          <LinkContainer className="mx-auto" to={`/signup`}>
            <Col as={Button} xs="12" variant="success">
              Create new account
            </Col>
          </LinkContainer>
        </Col>
      </Row>
    </Container>
  );
};

export default Login;
