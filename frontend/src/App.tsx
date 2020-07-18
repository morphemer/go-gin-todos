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

type DeleteUserID = IUser["ID"];

const App: React.FC = () => {
  const [users, setUsers] = useState<IUser[]>([]);
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");

  useEffect(() => {
    axios.get("http://localhost:8080/").then((res) => setUsers(res.data));
  }, []);

  const createUserHandler = async (e: FormEvent): Promise<void> => {
    try {
      e.preventDefault();

      const postParameter: postParameter = {
        username: username,
        email: email,
      };

      const response = await axios.post(
        "http://localhost:8080/users",
        postParameter
      );
      setUsers([...users, response.data]);
      setUsername("");
      setEmail("");
      console.log(`HTTP status ${response.status} OK`);
    } catch (error) {
      const { status, statusText, data } = error.response;
      alert(`HTTP status ${status} ${statusText} Message: ${data}`);
    }
  };

  const deleteUserHandler = (id: DeleteUserID): void => {
    axios
      .delete(`http://localhost:8080/users/${id}`)
      .then(() => {
        const filteredUsers = users.filter((user) => user.ID !== id);
        setUsers(filteredUsers);
        setUsername("");
        setEmail("");
        console.log(`user ${id} is deleted`);
      })
      .catch((error) => alert(error));
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
                  <th>Delete</th>
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
                      <td>
                        <Button
                          variant="danger"
                          onClick={() => deleteUserHandler(user.ID)}
                        >
                          Delete
                        </Button>
                      </td>
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
