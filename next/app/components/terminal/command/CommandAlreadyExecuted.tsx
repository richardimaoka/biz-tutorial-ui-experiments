interface Props {
  command: string;
}

export function CommandAlreadyExecuted(props: Props) {
  return <code>{props.command}</code>;
}
