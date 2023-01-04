import { css } from "@emotion/react";
import { Header } from "../components/Header";
import { IDE } from "../components/IDE";
import { Command, Terminal } from "../components/Terminal";

const list: Command[] = [
  {
    command: `docker build -t cmd1 -f Dockerfile.cmd1 .
docker run --rm cmd1`,
    output: "abc",
  },
  {
    command: `docker build -t cmd2 -f Dockerfile.cmd2 .
docker run --rm cmd2`,
    output: "abc def",
  },
  {
    command: `docker build -t cmd3 -f Dockerfile.cmd3 .
docker run --rm cmd3`,
    output: "abc",
  },
  {
    command: `docker build -t cmd4 -f Dockerfile.cmd4 .
docker run --rm cmd4`,
    output: "abc def",
  },
  {
    command: `docker build -t cmd5 -f Dockerfile.cmd5 .
docker run --rm cmd5`,
    output: `/home/your_username`,
  },
  {
    command: `docker build -t cmd6 -f Dockerfile.cmd6 .
docker run --rm cmd6`,
    output: `/home/your_username`,
  },
  {
    command: `docker pull nginx
docker inspect nginx`,
    output: `"Config": {
    "Cmd": [
        "nginx",
        "-g",
        "daemon off;"
    ]
}`,
  },
  {
    command: `docker run nginx:1.23.1`,
    output: `2022/09/03 07:14:58 [notice] 1#1: using the "epoll" event method
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
  },
  {
    command: `docker build -t cmd-nginx -f Dockerfile.cmd-nginx .
docker run --rm cmd-nginx`,
    output: `2022/09/03 05:48:46 [notice] 7#7: start worker process 13
2022/09/03 05:48:46 [notice] 7#7: start worker process 14
2022/09/03 05:48:46 [notice] 7#7: start worker process 15`,
  },
];

export default function Home() {
  return (
    <>
      <Header />
      <main
        css={css`
          background-color: #f8f8f8;
        `}
      >
        <div
          css={css`
            width: 680px;
            margin: 0 auto;
            background-color: white;
          `}
        >
          <IDE />
          <Terminal list={list} />
        </div>
      </main>
    </>
  );
}
