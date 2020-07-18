import React, { useState, useEffect, FormEvent } from "react";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Table from "react-bootstrap/Table";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { IUser } from "./models";
import axios from "axios";

interface postParameter {
  username: string;
  email: string;
}

const App: React.FC = () => {
  const [users, setUsers] = useState<IUser[]>([]);
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");

  useEffect(() => {
    axios.get("http://localhost:8080/").then((res) => setUsers(res.data));
  }, []);

  const createUserHandler = (e: FormEvent): void => {
    e.preventDefault();

    const postParameter: postParameter = {
      username: username,
      email: email,
    };

    axios
      .post("http://localhost:8080/users", postParameter)
      .then((res) => {
        setUsers([...users, res.data]);
        setUsername("");
        setEmail("");
      })
      .catch((error) => alert(error))
      .finally(() => {
        setUsername("");
        setEmail("");
      });
  };

  return (
    <div className="App">
      <Container>
        <Row className="justify-content-center">
          <Col>
            <h1>User list</h1>
            <Form>
              <Form.Group controlId="username">
                <Form.Control
                  type="text"
                  placeholder="Enter Username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </Form.Group>
              <Form.Group controlId="email">
                <Form.Control
                  type="email"
                  placeholder="Enter Email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </Form.Group>
              <Button
                variant="primary"
                type="submit"
                onClick={createUserHandler}
              >
                Submit
              </Button>
            </Form>
            <Table striped bordered hover>
              <thead>
                <tr>
                  <th>#</th>
                  <th>ID</th>
                  <th>UserName</th>
                  <th>Email</th>
                </tr>
              </thead>
              {users.map((user: IUser, idx: number) => {
                return (
                  <tbody key={idx}>
                    <tr>
                      <td>{idx}</td>
                      <td>{user.ID}</td>
                      <td>{user.Username}</td>
                      <td>{user.Email}</td>
                    </tr>
                  </tbody>
                );
              })}
            </Table>
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default App;
