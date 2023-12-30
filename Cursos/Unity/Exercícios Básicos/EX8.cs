using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class EX8 : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        int[,] matriz = {{1, 2}, {5, 3}};

        for(int i = 0; i < 2; i++)
            for(int j = 0; j < 2; j++)
                print("matriz["+ i + "," + j + "] = " + matriz[i,j]);
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
