// Copyright (c) 2023 Alyssia fung All rights reserved
//
// Created by: Alyssia fung
// Created on: Mar 2023
// This file contains the JS functions for index.html

/**
 * This function calculates area and perimeter of rectangle.
 */
"use strict"

function calculate() {
  // input
  const sideA = parseInt(document.getElementById("sideA").value)
  const sideB = parseInt(document.getElementById("sideB").value)
  const height = parseInt(document.getElementById("height").value)

  // process
  const area = [(A + B) / 2] * height

  // output
  document.getElementById("area").innerHTML = "Area is: " + area + " cm²"
}
