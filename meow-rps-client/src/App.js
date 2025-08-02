import logo from './logo.svg';
import './App.css';
import React, { useState } from 'react';
import styled from 'styled-components';

const Container = styled.div`
  font-family: 'Press Start 2P', cursive;
  text-align: center;
  padding: 40px;
  background-color: #fefefe;
`;

const Title = styled.h1`
  font-size: 20px;
  margin-bottom: 30px;
`;

const ButtonRow = styled.div`
  display: flex;
  justify-content: center;
  gap: 40px;
`;

const GameButton = styled.button`
  background-color: #fff;
  border: 4px solid #000;
  padding: 20px;
  font-size: 16px;
  cursor: pointer;
  font-family: inherit;
  image-rendering: pixelated;

  &:hover {
    background-color: #ddd;
  }
`;

const Result = styled.div`
  margin-top: 30px;
  font-size: 14px;
`;

const options = [
  { name: '가위', emoji: '✌️' },
  { name: '바위', emoji: '✊' },
  { name: '보', emoji: '✋' },
];

function App() {
  const [myPick, setMyPick] = useState(null);

  const handlePick = (pick) => {
    setMyPick(pick);
    // 여기에 WebSocket 전송 예정
  };

  return (
    <Container>
      <Title>🐾 MEOW RPS 🐾</Title>
      <ButtonRow>
        {options.map((opt) => (
          <GameButton key={opt.name} onClick={() => handlePick(opt.name)}>
            {opt.emoji}
            <br />
            {opt.name}
          </GameButton>
        ))}
      </ButtonRow>
      {myPick && <Result>선택한 손: {myPick}</Result>}
    </Container>
  );
}

export default App;