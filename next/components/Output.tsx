import { css } from "@emotion/react";

interface OutputProps {
  output: string;
}

export const Output = ({ output }: OutputProps): JSX.Element => {
  const cssTerminalCommand = css`
    margin: 1px 0px;
    padding: 4px;
    background-color: #3a3a3a;
    color: white;
    border-bottom: 1px solid white;
  `;
  return (
    <pre css={cssTerminalCommand}>
      <code>{output}</code>
    </pre>
  );
};
