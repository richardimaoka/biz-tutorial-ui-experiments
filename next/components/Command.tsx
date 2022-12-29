import { css } from "@emotion/react";

interface CommandProps {
  command: string;
  writtenLength: number;
}

export const Command = ({
  command,
  writtenLength,
}: CommandProps): JSX.Element => {
  const cssTerminalCommand = css`
    margin: 1px 0px;
    padding: 4px;
    background-color: #3a3a3a;
    color: white;
    border-bottom: 1px solid white;
  `;
  return (
    <pre css={cssTerminalCommand}>
      <code>{command.substring(0, writtenLength)}</code>
    </pre>
  );
};
