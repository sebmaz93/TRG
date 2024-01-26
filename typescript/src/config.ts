import path from "path";
import { Opts } from "./opts";

export enum Operation {
  Print,
  Add,
  Remove,
}

export type Config = {
  args: string[];
  operation: Operation;
  config: string;
  pwd: string;
};

function getPwd(opts: Opts): string {
  if (opts.pwd) {
    return opts.pwd;
  }

  return process.cwd();
}

function getConfig(opts: Opts): string {
  if (opts.config) {
    return opts.config;
  }

  const home = process.env["HOME"];
  const loc = process.env["XDG_CONFIG_HOME"] || home;
  if (!loc) {
    throw new Error("unable to determine config location");
  }

  if (loc === home) {
    return path.join(loc, ".projector.json");
  }
  return path.join(loc, "projector", "projector.json");
}

export default function config(opts: Opts): Config {
  return {
    pwd: getPwd(opts),
    config: getConfig(opts),
    args: getArgs(opts),
    operation: getOperation(opts),
  };
}
