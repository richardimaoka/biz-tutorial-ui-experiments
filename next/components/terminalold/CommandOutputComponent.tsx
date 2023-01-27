import { css } from "@emotion/react";

interface CommandOutputComponentProps {
  output: string;
}

export const CommandOutputComponent = ({
  output,
}: CommandOutputComponentProps): JSX.Element => {
  const cssTerminalCommand = css`
    margin: 1px 0px;
    padding: 4px;
    background-color: #1e1e1e;
    color: white;
    border-bottom: 1px solid #333333;
  `;
  return (
    <pre css={cssTerminalCommand}>
      <code>{output}</code>
    </pre>
  );
};
