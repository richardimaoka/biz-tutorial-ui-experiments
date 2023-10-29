import { TerminalComponent } from "@/app/components/terminal2/__TerminalComponent";

export default async function Page() {
  return (
    <div style={{ height: "100svh" }}>
      {/* <div
        style={{ width: "500px", height: "300px", backgroundColor: "white" }}
      >
        <div style={{ height: "100px", backgroundColor: "aqua" }}></div>
      </div> */}
      <TerminalComponent
        tabs={[
          { name: "default", href: "/test/terminal/" },
          { name: "another", href: "/test/terminal/" },
        ]}
        currentDirectory="/test/terminal"
        selectTab="default"
        entries={[
          {
            kind: "command",
            id: "cmd-1",
            command: "docker build -t cmd1 -f Dockerfile.cmd1 .",
            isExecuted: true,
          },
          {
            kind: "output",
            id: "output-1",
            output: `[+] Building 16.0s (6/6) FINISHED
 => [internal] load build definition from Dockerfile.cmd1                                                                                                                                                   0.0s
 => => transferring dockerfile: 80B                                                                                                                                                                         0.0s
 => [internal] load .dockerignore                                                                                                                                                                           0.0s
 => => transferring context: 2B                                                                                                                                                                             0.0s
 => [internal] load metadata for docker.io/library/ubuntu:latest                                                                                                                                            3.3s
 => [auth] library/ubuntu:pull token for registry-1.docker.io                                                                                                                                               0.0s
 => [1/1] FROM docker.io/library/ubuntu@sha256:aabed3296a3d45cede1dc866a24476c4d7e093aa806263c27ddaadbdce3c1054                                                                                            12.5s
 => => resolve docker.io/library/ubuntu@sha256:aabed3296a3d45cede1dc866a24476c4d7e093aa806263c27ddaadbdce3c1054                                                                                             0.0s
 => => sha256:445a6a12be2be54b4da18d7c77d4a41bc4746bc422f1f4325a60ff4fc7ea2e5d 29.54MB / 29.54MB                                                                                                           11.2s
 => => sha256:aabed3296a3d45cede1dc866a24476c4d7e093aa806263c27ddaadbdce3c1054 1.13kB / 1.13kB                                                                                                              0.0s
 => => sha256:b492494d8e0113c4ad3fe4528a4b5ff89faa5331f7d52c5c138196f69ce176a6 424B / 424B                                                                                                                  0.0s
 => => sha256:c6b84b685f35f1a5d63661f5d4aa662ad9b7ee4f4b8c394c022f25023c907b65 2.30kB / 2.30kB                                                                                                              0.0s
 => => extracting sha256:445a6a12be2be54b4da18d7c77d4a41bc4746bc422f1f4325a60ff4fc7ea2e5d                                                                                                                   0.9s
 => exporting to image                                                                                                                                                                                      0.0s
 => => exporting layers                                                                                                                                                                                     0.0s
 => => writing image sha256:710d7aba73162de2e4f1a7759908ddc6164ce9ea2deb0c7eccee358fdc701c16                                                                                                                0.0s
 => => naming to docker.io/library/cmd1`,
          },
        ]}
      />
    </div>
  );
}
