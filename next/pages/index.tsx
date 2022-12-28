import { useState } from "react";

const expected = [
  `docker build -t cmd-nginx -f Dockerfile.cmd-nginx .
docker run --rm cmd-nginx`,
  `2022/09/03 05:48:46 [notice] 7#7: start worker process 13
2022/09/03 05:48:46 [notice] 7#7: start worker process 14
2022/09/03 05:48:46 [notice] 7#7: start worker process 15
^C`,
];

export default function Home() {
  const [terminalElements, setTerminalElements] = useState<string[]>([]);
  const onClick = () => {
    if (terminalElements.length < expected.length) {
      console.log("append");
      const currentIndex = terminalElements.length - 1;
      const newElements = [...terminalElements];
      newElements.push(expected[currentIndex + 1]);
      setTerminalElements(newElements);
    }
  };
  return (
    <>
      {terminalElements.map((elem, index) => (
        <pre key={index}>
          <code>{elem}</code>
        </pre>
      ))}
      <button type="button" onClick={onClick}>
        run
      </button>
    </>
  );
}
