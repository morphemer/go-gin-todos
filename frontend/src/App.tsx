import React, { useState, useEffect } from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Table from 'react-bootstrap/Table'
import { IUser } from "./models"

const App: React.FC = () => {
  const [users, setUsers] = useState([])

  useEffect(() => {
    fetch("http://localhost:8080/")
      .then(res => res.json())
      .then(res => setUsers(res))
  }, []);

  console.log(users)

  return (
    <div className="App">
      <Container>
        <Row className="justify-content-center">
          <Col>
            <h1>User list</h1>
            <Table striped bordered hover>
              <thead>
                <tr>
                  <th>#</th>
                  <th>ID</th>
                  <th>UserName</th>
                  <th>Email</th>
                </tr>
              </thead>
              {
                users.map((user: IUser, idx: number) => {
                  return (
                    <tbody key={idx}>
                      <tr>
                        <td>{idx}</td>
                        <td>{user.ID}</td>
                        <td>{user.Username}</td>
                        <td>{user.Email}</td>
                      </tr>
                    </tbody>
                  )
                })
              }
            </Table>
          </Col>
        </Row>
      </Container>
    </div >
  );
}

export default App;
