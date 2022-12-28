import { css } from "@emotion/react";
import { useState } from "react";

const expected = [
  `docker build -t cmd1 -f Dockerfile.cmd1 .
docker run --rm cmd1`,
  "abc",
  `docker build -t cmd2 -f Dockerfile.cmd2 .
docker run --rm cmd2`,
  "abc def",
  `docker build -t cmd3 -f Dockerfile.cmd3 .
docker run --rm cmd3`,
  "abc",
  `docker build -t cmd4 -f Dockerfile.cmd4 .
docker run --rm cmd4`,
  "abc def",
  `docker build -t cmd5 -f Dockerfile.cmd5 .
docker run --rm cmd5`,
  `/home/your_username`,
  `docker build -t cmd6 -f Dockerfile.cmd6 .
docker run --rm cmd6`,
  `/home/your_username`,
  `docker pull nginx
docker inspect nginx`,
  `"Config": {
    "Cmd": [
        "nginx",
        "-g",
        "daemon off;"
    ]
}`,
  `docker run nginx:1.23.1`,
  `2022/09/03 07:14:58 [notice] 1#1: using the "epoll" event method
2022/09/03 07:14:58 [notice] 1#1: nginx/1.23.1
2022/09/03 07:14:58 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
2022/09/03 07:14:58 [notice] 1#1: OS: Linux 5.10.102.1-microsoft-standard-WSL2
2022/09/03 07:14:58 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2022/09/03 07:14:58 [notice] 1#1: start worker processes
2022/09/03 07:14:58 [notice] 1#1: start worker process 31
2022/09/03 07:14:58 [notice] 1#1: start worker process 32
2022/09/03 07:14:58 [notice] 1#1: start worker process 33
2022/09/03 07:14:58 [notice] 1#1: start worker process 34
2022/09/03 07:14:58 [notice] 1#1: start worker process 35
2022/09/03 07:14:58 [notice] 1#1: start worker process 36
2022/09/03 07:14:58 [notice] 1#1: start worker process 37
2022/09/03 07:14:58 [notice] 1#1: start worker process 38`,
  `docker build -t cmd-nginx -f Dockerfile.cmd-nginx .
docker run --rm cmd-nginx`,
  `2022/09/03 05:48:46 [notice] 7#7: start worker process 13
2022/09/03 05:48:46 [notice] 7#7: start worker process 14
2022/09/03 05:48:46 [notice] 7#7: start worker process 15`,
  `^C`,
];

export const Terminal = (): JSX.Element => {
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
    <div>
      <div
        css={css`
          background-color: #3a3a3a;
          height: 300px;
        `}
      >
        {terminalElements.map((elem, index) => (
          <pre
            css={css`
              margin: 1px 0px;
              padding: 4px;
              background-color: #3a3a3a;
              color: white;
              border-bottom: 1px solid white;
            `}
            key={index}
          >
            <code>{elem}</code>
          </pre>
        ))}
      </div>

      <button type="button" onClick={onClick}>
        run
      </button>
    </div>
  );
};
