import { FlickeringTrail } from "./FlickeringTrail";

interface Props {
  command: string;
}

export function CommandWaitingExecution(props: Props) {
  return (
    <code>
      {props.command}
      <FlickeringTrail />
    </code>
  );
}
