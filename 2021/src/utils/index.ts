/**
 * Root for your util libraries.
 *
 * You can import them in the src/template/index.ts,
 * or in the specific file.
 *
 * Note that this repo uses ES Modules, so you have to explicitly specify
 * .js extension (yes, .js not .ts - even for TypeScript files)
 * for imports that are not imported from node_modules.
 *
 * For example:
 *
 *   correct:
 *
 *     import _ from 'lodash'
 *     import myLib from '../utils/myLib.js'
 *     import { myUtil } from '../utils/index.js'
 *
 *   incorrect:
 *
 *     import _ from 'lodash'
 *     import myLib from '../utils/myLib.ts'
 *     import { myUtil } from '../utils/index.ts'
 *
 *   also incorrect:
 *
 *     import _ from 'lodash'
 *     import myLib from '../utils/myLib'
 *     import { myUtil } from '../utils'
 *
 */

import fs from "fs";
import path from "path";

export const parseInput = (rawInput: string) =>
  rawInput
    .lines()
    .map((line) => line.trim())
    .join("\n");

export const parseBinary = (string: string) => parseInt(string, 2);
export const parseDecimal = (string: string) => parseInt(string, 10);

declare global {
  interface Array<T> {
    first(): T;
    reduceWhile<U>(
      callbackfn: (
        previousValue: U,
        currentValue: T,
        currentIndex: number,
        array: T[]
      ) => U,
      predicate: (
        previousValue: U,
        currentValue: T,
        currentIndex: number,
        array: T[]
      ) => boolean,
      initialValue?: U
    ): U;
    tupleWindows(size: number): Array<Array<T>>;
  }

  interface Object {
    map<T, U>(this: T, callbackfn: (value: T) => U): U;
  }

  interface String {
    chars(): string[];
    lines(): string[];
  }
}

Array.prototype.first = function () {
  return this[0];
};

Array.prototype.reduceWhile = function <T, U>(
  callbackfn: (
    previousValue: U,
    currentValue: T,
    currentIndex: number,
    array: T[]
  ) => U,
  predicate: (
    previousValue: U,
    currentValue: T,
    currentIndex: number,
    array: T[]
  ) => boolean,
  initialValue?: U
) {
  return this.reduce((previousValue, currentValue, currentIndex, array) => {
    if (predicate(previousValue, currentValue, currentIndex, array)) {
      return previousValue;
    } else {
      return callbackfn(previousValue, currentValue, currentIndex, array);
    }
  }, initialValue);
};

Array.prototype.tupleWindows = function (size: number) {
  if (this.length <= size) {
    return [this];
  }

  return this.reduce((result, _, i, values) => {
    if (i < values.length - size + 1) {
      return [
        ...result,
        Array.from({ length: size }).map((_, n) => values[i + n]),
      ];
    }

    return result;
  }, []);
};

Object.prototype.map = function <T, U>(
  this: T,
  callbackFn: (value: T) => U
): U {
  return callbackFn(this);
};

String.prototype.chars = function () {
  return this.split("");
};

String.prototype.lines = function () {
  return this.split("\n");
};
