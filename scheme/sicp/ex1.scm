(define (f n)
  (cond ((< n 3) (f n))
        (
         (> n 3) (+ (f (- n 1)) (* 2 (f (- n 2))) (* 3 (f (- n 3))))
        )
