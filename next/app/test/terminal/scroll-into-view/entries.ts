import { TerminalEntry } from "@/app/components/terminal2/definitions";
export const entries: TerminalEntry[] = [
  {
    id: "66105871-d23c-4e54-98f4-b715b330a0da",
    kind: "command",
    command: "mkdir myproj",
    isExecuted: true,
  },
  {
    id: "68b0d161-d61d-45b2-bf8a-e7842a59375d",
    kind: "command",
    command: "cd myproj",
    isExecuted: true,
  },
  {
    id: "1fd65263-a163-4856-9e8e-22846794bf30",
    kind: "command",
    command: "touch Dockerfile.cmd1",
    isExecuted: true,
  },
  {
    id: "ffc99a30-24c2-4ed5-9c2f-987a37007181",
    kind: "command",
    command: "docker build -t cmd1 -f Dockerfile.cmd1 .",
    isExecuted: true,
  },
  {
    id: "29af3a1b-af1c-4472-857f-ff55271dccc2",
    kind: "output",
    output:
      "[+] Building 16.0s (6/6) FINISHED\n => [internal] load build definition from Dockerfile.cmd1                                                                                                                                                   0.0s\n => => transferring dockerfile: 80B                                                                                                                                                                         0.0s\n => [internal] load .dockerignore                                                                                                                                                                           0.0s\n => => transferring context: 2B                                                                                                                                                                             0.0s\n => [internal] load metadata for docker.io/library/ubuntu:latest                                                                                                                                            3.3s\n => [auth] library/ubuntu:pull token for registry-1.docker.io                                                                                                                                               0.0s\n => [1/1] FROM docker.io/library/ubuntu@sha256:aabed3296a3d45cede1dc866a24476c4d7e093aa806263c27ddaadbdce3c1054                                                                                            12.5s\n => => resolve docker.io/library/ubuntu@sha256:aabed3296a3d45cede1dc866a24476c4d7e093aa806263c27ddaadbdce3c1054                                                                                             0.0s\n => => sha256:445a6a12be2be54b4da18d7c77d4a41bc4746bc422f1f4325a60ff4fc7ea2e5d 29.54MB / 29.54MB                                                                                                           11.2s\n => => sha256:aabed3296a3d45cede1dc866a24476c4d7e093aa806263c27ddaadbdce3c1054 1.13kB / 1.13kB                                                                                                              0.0s\n => => sha256:b492494d8e0113c4ad3fe4528a4b5ff89faa5331f7d52c5c138196f69ce176a6 424B / 424B                                                                                                                  0.0s\n => => sha256:c6b84b685f35f1a5d63661f5d4aa662ad9b7ee4f4b8c394c022f25023c907b65 2.30kB / 2.30kB                                                                                                              0.0s\n => => extracting sha256:445a6a12be2be54b4da18d7c77d4a41bc4746bc422f1f4325a60ff4fc7ea2e5d                                                                                                                   0.9s\n => exporting to image                                                                                                                                                                                      0.0s\n => => exporting layers                                                                                                                                                                                     0.0s\n => => writing image sha256:710d7aba73162de2e4f1a7759908ddc6164ce9ea2deb0c7eccee358fdc701c16                                                                                                                0.0s\n => => naming to docker.io/library/cmd1",
  },
  {
    id: "a431d2bf-232d-4406-8f1c-f8606f719a88",
    kind: "command",
    command: "docker run --rm cmd1",
    isExecuted: true,
  },
  {
    id: "f42cd3bb-b1c9-442a-9baf-381cc3247657",
    kind: "output",
    output: "abc",
  },
  {
    id: "effb5cd5-5936-45b9-b505-13240469d1ff",
    kind: "command",
    command: "cp Dockerfile.cmd1 Dockerfile.cmd2",
    isExecuted: true,
  },
  {
    id: "4ca88cff-4b85-445b-8dd2-c2fecbb9d9a1",
    kind: "command",
    command: "docker build -t cmd2 -f Dockerfile.cmd2 .",
    isExecuted: true,
  },
  {
    id: "4c951ebb-d0cf-4031-a054-cbc0b8e27ab2",
    kind: "output",
    output:
      "[+] Building 1.1s (5/5) FINISHED\n => [internal] load build definition from Dockerfile.cmd2                                                                                                                                                   0.0s\n => => transferring dockerfile: 87B                                                                                                                                                                         0.0s\n => [internal] load .dockerignore                                                                                                                                                                           0.0s\n => => transferring context: 2B                                                                                                                                                                             0.0s\n => [internal] load metadata for docker.io/library/ubuntu:latest                                                                                                                                            0.9s\n => CACHED [1/1] FROM docker.io/library/ubuntu@sha256:aabed3296a3d45cede1dc866a24476c4d7e093aa806263c27ddaadbdce3c1054                                                                                      0.0s\n => exporting to image                                                                                                                                                                                      0.0s\n => => exporting layers                                                                                                                                                                                     0.0s\n => => writing image sha256:14a01cce2db6d56f25f522954ba0aaabc66478f1ecde30336ab1f3e83288a3ac                                                                                                                0.0s\n => => naming to docker.io/library/cmd2",
  },
  {
    id: "18cde1b6-c189-48ad-9c7b-709dd29de77e",
    kind: "command",
    command: "docker run --rm cmd2",
    isExecuted: true,
  },
  {
    id: "8605315f-9fc1-4d84-9528-eae35436fc13",
    kind: "command",
    command: "docker pull httpd",
    isExecuted: true,
  },
  {
    id: "9b2d65fb-6958-427e-bed2-21fb4828566b",
    kind: "command",
    command: "docker inspect httpd",
    isExecuted: true,
  },
  {
    id: "e07ee936-2fb2-4103-8b51-47b61c490896",
    kind: "command",
    command: "docker run httpd",
    isExecuted: true,
  },
  {
    id: "f62c9b46-683d-4532-817f-49f569d8fa53",
    kind: "command",
    command: "docker pull python",
    isExecuted: true,
  },
  {
    id: "3c661342-9dc0-4b1b-a7da-1853489a5763",
    kind: "command",
    command: "docker inspect python",
    isExecuted: true,
  },
  {
    id: "2e7de957-312c-420a-9c0b-f4af2a786b1e",
    kind: "command",
    command: "docker run python",
    isExecuted: true,
  },
  {
    id: "4cdb47b0-7557-48db-93d7-710890223987",
    kind: "output",
    output: "",
  },
  {
    id: "06ab45f1-e594-44e1-846c-9dc362406db6",
    kind: "command",
    command: "docker run -it python",
    isExecuted: true,
  },
  {
    id: "35b4fe2c-4612-4e51-af5e-fdf34a018cee",
    kind: "output",
    output:
      'Python 3.12.0 (main, Oct  3 2023, 01:48:15) [GCC 12.2.0] on linux\nType "help", "copyright", "credits" or "license" for more information',
  },
  {
    id: "a886d29a-0f06-41a0-8044-c72b98df74ff",
    kind: "output",
    output: '>>> print("hello python")',
  },
  {
    id: "8e21b6fd-84e5-40af-acb1-e498ad80e19e",
    kind: "output",
    output: "hello python\n>>>",
  },
  {
    id: "22f671d0-a908-4fb9-9b2f-8774ba918ab8",
    kind: "output",
    output: "(Ctrl + D)",
  },
  {
    id: "3f0b0d48-9d22-4bb7-b98c-6653c8575a57",
    kind: "command",
    command: "docker pull nginx",
    isExecuted: true,
  },
  {
    id: "f378cbe5-d1d0-4e62-99b8-84e9b0b83e3a",
    kind: "command",
    command: "docker inspect nginx",
    isExecuted: true,
  },
  {
    id: "aaaa",
    kind: "command",
    command: "docker run nginx",
    isExecuted: true,
  },
];
